func findLeadersNaive(arr []int) []int {
	n := len(arr)
	leaders := []int{}

	for i := 0; i < n; i++ {
		isLeader := true
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[i] {
				isLeader = false
				break
			}
		}
		if isLeader {
			leaders = append(leaders, arr[i])
		}
	}
	return leaders
}
