package eventBuckets

func (fb *Bucket) Count() uint64 {
	return fb.totalCount
}
