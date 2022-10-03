package creationalpatterns

// My notes: The code in below.
// creational patterns try to give ready-to-use objects to users instead of asking for their creation
//single instance of an object, and guarantee that there are no duplicates

// At the first call to use the instance, it is created and then reused between all the parts in the application that need to use that particular behavior

// When you want to use the same connection to a database to make every query (Or save data from config's file to a config struct)'

// When you open a Secure Shell (SSH) connection to a server to do a few tasks, and don't want to reopen the connection for each task
// If you need to limit the access to some variable or space, you use a Singleton as the door to this variable (we'll see in the following chapters that this is more achievable in Go using channels anyway)
// If you need to limit the number of calls to some places, you create a Singleton instance to make the calls in the accepted window

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var Instance *singleton

func GetIntances() Singleton {
	if Instance == nil {
		Instance = new(singleton)
	}
	return Instance
}

func (s *singleton) AddOne() int {
	s.count++
       return s.count
}
