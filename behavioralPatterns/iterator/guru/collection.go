package behavioralPatterns

type Collection interface {
    createIterator() Iterator
}