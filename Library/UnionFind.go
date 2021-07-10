package UnionFind

/*
n: 要素数
parent[i]: iの親となる要素
siz[i]: iを含む素集合のサイズ
group[i]j: 代表元をiとする素集合に含まれる要素の配列j
*/
type UnionFind struct {
	n      int
	parent []int
	siz    []int
	group  map[int][]int
}

/*
コンストラクタ
任意のiについて、parent[i]=i、siz[i]=1、として初期化
group[i]jは空のmapとして初期化
*/
func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.n = n
	uf.parent = make([]int, n)
	uf.siz = make([]int, n)
	uf.group = make(map[int][]int)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.siz[i] = 1
	}
	return uf
}

/*
要素xの属する木（素集合）の根の値を取得し、インデックスを返す
同時に経路圧縮も行う
*/
func (uf *UnionFind) root(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		uf.parent[x] = uf.root(uf.parent[x])
		return uf.parent[x]
	}
}

/*
要素xと要素yを含む素集合を併合する
*/
func (uf *UnionFind) unite(x, y int) {
	rx, ry := uf.root(x), uf.root(y)
	if rx == ry {
		return
	}
	if uf.siz[rx] < uf.siz[ry] {
		rx, ry = ry, rx
	}
	uf.siz[rx] += uf.siz[ry]
	uf.parent[ry] = rx
	return
}

/*
要素xと要素yが同じ素集合に属するかを判定し、真偽値を返す
*/
func (uf *UnionFind) same(x, y int) bool {
	rx, ry := uf.root(x), uf.root(y)
	return rx == ry
}

func (uf *UnionFind) size(x int) int {
	return uf.siz[uf.root(x)]
}

/*
現在の素集合の状態に基づいてgroupを更新する
O(uf.n)なので注意
*/
func (uf *UnionFind) grouping() {
	for i := 0; i < uf.n; i++ {
		uf.root(i)
	}
	for i := 0; i < uf.n; i++ {
		uf.group[uf.parent[i]] = append(uf.group[uf.parent[i]], i)
	}
}

/*
素集合を初期化する
*/
func (uf *UnionFind) clear() {
	for i := 0; i < uf.n; i++ {
		uf.parent[i] = i
		uf.siz = make([]int, uf.n)
	}
	uf.group = make(map[int][]int)
}
