package words

// 这个文件是用来获取单词的，获取单词的描述的，以及关闭数据库连接的

type Word struct {
	Id          int
	Word        string
	Discription []string
	Src         string
}

type WordList struct {
	Words []Word
}

func NewWordList() *WordList {
	return &WordList{
		Words: []Word{},
	}
}

func (wl *WordList) GetWords(db *Database, query string) {
	rows, err := db.Db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var word Word
		// 按照元组的顺序扫描数据
		if err := rows.Scan(&word.Id, &word.Word, &word.Src); err != nil {
			panic(err)
		}
		wl.Words = append(wl.Words, word)
	}
}

func (wl *WordList) GetWordsArray() []Word {
	return wl.Words
}

func (wl *WordList) GetWordsDiscription(db *Database) {

	query := "SELECT word_id,translation FROM translations"
	rows, err := db.Db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var wordId int
		var translation string
		err = rows.Scan(&wordId, &translation)
		if err != nil {
			panic(err)
		}
		for i, word := range wl.Words {
			if word.Id == wordId {
				wl.Words[i].Discription = append(wl.Words[i].Discription, translation)
				break
			}
		}
	}

}
