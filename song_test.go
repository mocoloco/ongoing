package song_test

import (
	"flag"
	"fmt"
	"io"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"

	pb "."
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	number_of_client = flag.Int("number_of_client", 1, "number of client to start")
	port             = flag.Int("port", 10000, "The server port")
	serverAddr       = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

//!!!!!!!!!!!!!Server API!!!!!!!!!!!!

type SongSrvServer struct {
	savedSongs []pb.SongObj
}

// generate new server
func newServer() *SongSrvServer {

	s := new(SongSrvServer)
	return s
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func match(s1, s2 string) bool {
	return s1 == s2
}

func regex(s, re string) bool {
	var validSong = regexp.MustCompile(re)
	return validSong.MatchString(s)
}

// Add: add one song to DB.
func (s *SongSrvServer) Add(ctx context.Context, song_obj *pb.SongObj) (*pb.SongResponce, error) {

	grpclog.Printf("SERVER: Add song request with Id %v\n", song_obj.Id)
	s.savedSongs = append(s.savedSongs, *song_obj)
	return &pb.SongResponce{song_obj.Id}, nil
}

// Adds: add one or more songs to DB.
func (s *SongSrvServer) Adds(stream pb.SongSrv_AddsServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		s.savedSongs = append(s.savedSongs, *in)
		if err := stream.Send(&pb.SongResponce{in.Id}); err != nil {
			return err

		}
	}
}

// Get: get one song from DB.
func (s *SongSrvServer) Get(ctx context.Context, SongObj *pb.SongObj) (*pb.SongObj, error) {

	grpclog.Printf("SERVER: Recive Get Request for Song Id %v\n", SongObj.Id)
	for _, song := range s.savedSongs {
		if song.Id == SongObj.Id {
			return &song, nil
		}
	}
	// No SongObj was found, return an unnamed feature
	return nil, fmt.Errorf("No such song id: %v", SongObj.Id)
}

// Modify: modify one song from DB.
func (s *SongSrvServer) Modify(ctx context.Context, SongObj *pb.SongObj) (*pb.SongResponce, error) {

	grpclog.Printf("SERVER: Recive Modify Request for Song Id %v\n", SongObj.Id)
	for i, song := range s.savedSongs {
		if song.Id == SongObj.Id {
			s.savedSongs = append(s.savedSongs[:i], s.savedSongs[i+1:]...)
			s.savedSongs = append(s.savedSongs, *SongObj)
			return &pb.SongResponce{SongObj.Id}, nil
		}
	}
	// No SongObj was found, return an unnamed feature
	return nil, fmt.Errorf("No such song id: %v", SongObj.Id)
}

// Delete: delete one song from DB.
func (s *SongSrvServer) Delete(ctx context.Context, SongObj *pb.SongObj) (*pb.SongResponce, error) {

	grpclog.Printf("SERVER: Recive Delete Request for Song Id %v\n", SongObj.Id)
	for i, song := range s.savedSongs {
		if song.Id == SongObj.Id {
			s.savedSongs = append(s.savedSongs[:i], s.savedSongs[i+1:]...)
			return &pb.SongResponce{SongObj.Id}, nil
		}
	}
	// No SongObj was found, return an unnamed feature
	return nil, fmt.Errorf("No such song id: %v", SongObj.Id)
}

func (s *SongSrvServer) Filter(query *pb.SongQuery, stream pb.SongSrv_FilterServer) error {

	grpclog.Printf("SERVER: Recive Filter Request : %v\n", query)
	var triger func(string, string) bool
	if query.SearchType == 0 {
		triger = regex
	} else if query.SearchType == 2 {
		triger = contains
	} else if query.SearchType == 3 {
		triger = match
	}

	for _, song := range s.savedSongs {
		var str string
		if query.SearchField == 0 {
			str = song.Tags.Title
		} else if query.SearchField == 1 {
			str = song.Tags.Artist
		} else if query.SearchField == 2 {
			str = song.Tags.Album
		}
		if triger(str, query.Search) {
			if err := stream.Send(&song); err != nil {
				return err
			}
		}
	}
	return nil
}

//!!!!!!!!!!Client Api !!!!!!!!!!!!!!!!!

// return uniq id for each SongObj
func uniq_id(SongsObj []pb.SongObj, client_id int) []pb.SongObj {

	// Use _SongsObj for retrive list of SongObj with uniq ID
	// Use []pb.SongObj and not []*pb.SongObj because I want each client
	// will have a Uniq and readable ID
	var _SongsObj []pb.SongObj
	for _, obj := range SongsObj {
		obj.Id = fmt.Sprintf("%v__%v", strconv.Itoa(client_id), obj.Id)
		_SongsObj = append(_SongsObj, obj)
	}
	return _SongsObj
}

func motivate(action string, list_of_chan []chan map[string]interface{}, wg *sync.WaitGroup, songs interface{}) {

	for _, ch := range list_of_chan {
		wg.Add(1)
		_map := make(map[string]interface{})
		_map[action] = songs
		ch <- _map
	}
	wg.Wait()

}

// add_song add one song to DB.
func add_song(client pb.SongSrvClient, SongsObj []pb.SongObj, client_id int) {

	SongObj := uniq_id(SongsObj, client_id)[0]
	grpclog.Printf("CLIENT-%v: Call to Add Function with Title %v and Id %v\n", client_id, SongObj.Tags.Title, SongObj.Id)
	status, err := client.Add(context.Background(), &SongObj)
	if err != nil {
		grpclog.Fatalf("CLIENT-%v: %v.add_song(_) = _, %v: ", client_id, client, err)
	}
	grpclog.Printf("CLIENT-%v: Got Response from Server that Song id %v was Add succesfully to DB\n\n", client_id, status.Id)
}

// add_songs add one or more songs to DB.
func add_songs(client pb.SongSrvClient, SongsObj []pb.SongObj, client_id int) {
	grpclog.Printf("CLIENT-%v: Call to Adds Function with list of %v SongObj\n", client_id, len(SongsObj))
	songs := uniq_id(SongsObj, client_id)
	stream, err := client.Adds(context.Background())
	if err != nil {
		grpclog.Fatalf("%v.add_songs(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				grpclog.Fatalf("CLIENT-%v: Failed to add song Id %v: %v", client_id, in.Id, err)
			}
			grpclog.Printf("CLIENT-%v: Successful add song id %v\n", client_id, in.Id)
		}
	}()
	for _, song := range songs {
		if err := stream.Send(&song); err != nil {
			grpclog.Fatalf("CLIENT-%v: Failed to send a song id %v : %v", client_id, song.Id, err)
		}
	}
	stream.CloseSend()
	<-waitc
	grpclog.Printf("\n")
}

// get_song get one song by id.
func get_song(client pb.SongSrvClient, SongsObj []pb.SongObj, client_id int) {

	SongObj := uniq_id(SongsObj, client_id)[0]
	grpclog.Printf("CLIENT-%v: Send Get Request for Song Id %v\n", client_id, SongObj.Id)
	song, err := client.Get(context.Background(), &SongObj)
	if err != nil {
		grpclog.Printf("CLIENT-%v: %v.get_song(_) = _, %v: \n", client_id, client, err)
	} else {
		grpclog.Printf("CLIENT-%v: Successful Get Song Obj %v\n\n", client_id, song)
	}
}

// modify_song modify one song by id.
func modify_song(client pb.SongSrvClient, SongsObj []pb.SongObj, client_id int) {

	SongObj := uniq_id(SongsObj, client_id)[0]
	grpclog.Printf("CLIENT-%v: Request Modify Song Id %v\n", client_id, SongObj.Id)
	status, err := client.Modify(context.Background(), &SongObj)
	if err != nil {
		grpclog.Fatalf("%v.modify_song(_) = _, %v: ", client, err)
	}
	grpclog.Printf("CLIENT-%v: Successful Modify Song Obj %v\n\n", client_id, status)
}

// delete_song modify one song by id.
func delete_song(client pb.SongSrvClient, SongsObj []pb.SongObj, client_id int) {

	SongObj := uniq_id(SongsObj, client_id)[0]
	grpclog.Printf("CLIENT-%v: Request Delete Song Id %v\n", client_id, SongObj.Id)
	status, err := client.Delete(context.Background(), &SongObj)
	if err != nil {
		grpclog.Fatalf("%v.modify_song(_) = _, %v: ", client, err)
	}
	grpclog.Printf("CLIENT-%v: Successful Delete Song Obj %v\n\n", client_id, status)
}

// filter_songs lists all the song within the given query.
func filter_songs(client pb.SongSrvClient, query pb.SongQuery, client_id int) {

	grpclog.Printf("CLIENT-%v: Request Filter Song Id %v\n", client_id, query)
	stream, err := client.Filter(context.Background(), &query)
	if err != nil {
		grpclog.Fatalf("CLIENT-%v: %v.filter_songs(_) = _, %v", client_id, client, err)
	}
	for {
		song, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Fatalf("CLIENT-%v: %v.filter_songs(_) = _, %v", client_id, client, err)
		}
		grpclog.Printf("CLIENT-%v: Get Song Obj %v\n", client_id, song)
	}
	grpclog.Printf("\n")
}

func start_client(client_id int, map_cahn chan map[string]interface{}, wg *sync.WaitGroup) {

	var dial_opts []grpc.DialOption
	dial_opts = append(dial_opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, dial_opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSongSrvClient(conn)
	grpclog.Printf("CLIENT-%v: Successfully Connect\n", client_id)
	for {
		select {
		case action := <-map_cahn:
			if SongObj, ok := action["add_song"]; ok {
				add_song(client, SongObj.([]pb.SongObj), client_id)
				wg.Done()
			} else if SongsObj, ok := action["add_songs"]; ok {
				add_songs(client, SongsObj.([]pb.SongObj), client_id)
				wg.Done()
			} else if SongObj, ok := action["get_song"]; ok {
				get_song(client, SongObj.([]pb.SongObj), client_id)
				wg.Done()
			} else if SongObj, ok := action["modify_song"]; ok {
				modify_song(client, SongObj.([]pb.SongObj), client_id)
				wg.Done()
			} else if SongObj, ok := action["delete_song"]; ok {
				delete_song(client, SongObj.([]pb.SongObj), client_id)
				wg.Done()
			} else if SongQuery, ok := action["filter_songs"]; ok {
				filter_songs(client, SongQuery.(pb.SongQuery), client_id)
				wg.Done()
			}
		}
	}

}

func TestMain(m *testing.M) {

	flag.Parse()
	var wg sync.WaitGroup
	var srv_opts []grpc.ServerOption
	var songs []pb.SongObj

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(srv_opts...)
	pb.RegisterSongSrvServer(grpcServer, newServer())
	grpclog.Println("SERVER: Successfully bind")
	go func() {
		grpcServer.Serve(lis)
	}()

	// Start Client that wait for  action
	list_of_chan := make([]chan map[string]interface{}, 0)
	for i := 0; i < *number_of_client; i++ {
		map_chan := make(chan map[string]interface{})
		list_of_chan = append(list_of_chan, map_chan)
		go start_client(i, map_chan, &wg)
	}

	// send add_song and SongObj to client channel
	grpclog.Println("MAIN: CALL TO add_song with one SongObj")
	songs = []pb.SongObj{
		{
			Tags: &pb.Tags{Title: "Lazarus", Artist: "David Bowie", Album: "Herors"}, Id: "1",
		},
	}
	motivate("add_song", list_of_chan, &wg, songs)

	// send add_songs and 3 SongObj to client channel
	grpclog.Println("MAIN: CALL TO add_songs with 3 SongObj")
	songs = []pb.SongObj{
		{
			Tags: &pb.Tags{Title: "I'm Lost", Artist: "KNOB", Album: "Legend"}, Id: "2",
		},
		{
			Tags: &pb.Tags{Title: "Don't Tell Me", Artist: "KNOB", Album: "Legend"}, Id: "3",
		},
		{
			Tags: &pb.Tags{Title: "Clocks", Artist: "Coldplay", Album: "Magic"}, Id: "4",
		},
	}
	motivate("add_songs", list_of_chan, &wg, songs)

	// send get_song with 1 SongObj id to client channel
	grpclog.Println("MAIN: CALL TO get_song with 1 SongObj")
	motivate("get_song", list_of_chan, &wg, []pb.SongObj{{Id: "1"}})

	// send modify_song and 1 SongObj to client channel
	songs = []pb.SongObj{
		{
			Tags: &pb.Tags{Title: "mocoloco", Artist: "David Bowie", Album: "Herors"}, Id: "1",
		},
	}
	grpclog.Println("MAIN: CALL TO modify_song with 1 SongObj")
	motivate("modify_song", list_of_chan, &wg, songs)

	// send get_song with 1 SongObj id to client channel
	grpclog.Println("MAIN: CALL TO get_song with 1 SongObj")
	motivate("get_song", list_of_chan, &wg, []pb.SongObj{{Id: "1"}})

	// send filter_songs with 1 SongObj id to client channel
	grpclog.Println("MAIN: CALL TO filter_songs with SongQuery")
	motivate("filter_songs", list_of_chan, &wg, pb.SongQuery{Search: "KN", SearchType: 2, SearchField: 1})

	// send get_song with 1 SongObj id to client channel
	grpclog.Println("MAIN: CALL TO delete_song with 1 SongObj")
	motivate("delete_song", list_of_chan, &wg, []pb.SongObj{{Id: "1"}})

	// send get_song with 1 SongObj id to client channel
	grpclog.Println("MAIN: CALL TO get_song with 1 SongObj")
	motivate("get_song", list_of_chan, &wg, []pb.SongObj{{Id: "1"}})

}
