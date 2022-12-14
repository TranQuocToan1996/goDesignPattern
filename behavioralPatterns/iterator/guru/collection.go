package behavioralPatterns

type Collection interface {
    CreateIterator() Iterator
}