package uerrors

/*
Collects errors from the given error channel into an aggregate of errors.
The caller must close errSrc when no more errors are expected.
In turn, a single aggregate error will be placed on the output channel,
followed by closing that channel.
*/
func CollectChan[E error](errSrc <-chan E) <-chan *LinkedAggregate {
	collectedErrs := make(chan *LinkedAggregate)
	go func() {
		aggErrs := LinkedAggregate{}
		defer func() {
			collectedErrs <- &aggErrs
			close(collectedErrs)
		}()
		for err := range errSrc {
			aggErrs.Add(err)
		}
	}()
	return collectedErrs
}
