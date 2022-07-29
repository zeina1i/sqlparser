package sqlparser

type (
	Column struct {
	}

	Expression struct {
	}

	CreateTable struct {
		name    string
		columns []Column
	}

	Delete struct {
		table string
		where Expression
	}

	Insert struct {
		table   string
		columns []string
		values  []string
	}

	Update struct {
		table   string
		columns []string
		values  []string
		where   Expression
	}

	Select struct {
		table  string
		fields []string
		where  Expression
	}
)
