// https://leetcode.com/problems/all-paths-from-source-to-target/?envType=study-plan&id=algorithm-ii
package main

func main(){

}
func dfs(graph [][]int, node int, paths *[][]int, path []int) {
    if node == len(graph)-1 {
        *paths = append(*paths, path)
        return
    }

    for _, next := range graph[node] {
        newPath := make([]int, len(path))
        copy(newPath, path)
        newPath = append(newPath, next)
        dfs(graph, next, paths, newPath)
    }
}
func allPathsSourceTarget(graph [][]int) [][]int {
    var paths [][]int
    var path []int
    path = append(path, 0)

    dfs(graph, 0, &paths, path)

    return paths
}
