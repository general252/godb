package model
// 此文件是根据GoGoModelsField自动生成


    // Book database table name is "books"
    type GoBook struct {
    
        mID string // 
    
        mUid string // gorm:"column:uid;type:string;uniqueIndex;not null"
    
        mCreatedAt string // 
    
        mUpdatedAt string // 
    
        mAuthor string // gorm:"column:author;type:string;default:'';size:64"
    
        mName string // gorm:"column:name;type:string;default:'';size:128"
    

        cols []string // 字段列表
    }

    func newGoBook() *GoBook {
        return &GoBook {
            
                mID: "id",
            
                mUid: "uid",
            
                mCreatedAt: "created_at",
            
                mUpdatedAt: "updated_at",
            
                mAuthor: "author",
            
                mName: "name",
            
        }
    }

    func (*Book) TableName() string {
    return "books"
    }

    // 函数
    
        func (c *GoBook) FieldNameID() string {
            return c.mID
        }
    
        func (c *GoBook) FieldNameUid() string {
            return c.mUid
        }
    
        func (c *GoBook) FieldNameCreatedAt() string {
            return c.mCreatedAt
        }
    
        func (c *GoBook) FieldNameUpdatedAt() string {
            return c.mUpdatedAt
        }
    
        func (c *GoBook) FieldNameAuthor() string {
            return c.mAuthor
        }
    
        func (c *GoBook) FieldNameName() string {
            return c.mName
        }
    

    // 字段
    
        func (c *GoBook) AddColID() *GoBook {
            c.cols = append(c.cols, c.mID)
            return c
        }
    
        func (c *GoBook) AddColUid() *GoBook {
            c.cols = append(c.cols, c.mUid)
            return c
        }
    
        func (c *GoBook) AddColCreatedAt() *GoBook {
            c.cols = append(c.cols, c.mCreatedAt)
            return c
        }
    
        func (c *GoBook) AddColUpdatedAt() *GoBook {
            c.cols = append(c.cols, c.mUpdatedAt)
            return c
        }
    
        func (c *GoBook) AddColAuthor() *GoBook {
            c.cols = append(c.cols, c.mAuthor)
            return c
        }
    
        func (c *GoBook) AddColName() *GoBook {
            c.cols = append(c.cols, c.mName)
            return c
        }
    

    func (c *GoBook) AllCols() []string {
        return c.cols
    }

    func (c *GoBook) ResetCols() {
        c.cols = []string{}
    }

    func (c *GoBook) AddAllCols() {
        c.cols = append(c.cols, "*")
    }

    // User database table name is "users"
    type GoUser struct {
    
        mID string // 
    
        mUid string // gorm:"column:uid;type:string;uniqueIndex;not null"
    
        mCreatedAt string // 
    
        mUpdatedAt string // 
    
        mName string // 
    
        mAge string // 
    
        mBirthday string // 
    
        mCompanyID string // 
    
        mManagerID string // gorm:"column:manager_id;type:int;default:0"
    

        cols []string // 字段列表
    }

    func newGoUser() *GoUser {
        return &GoUser {
            
                mID: "id",
            
                mUid: "uid",
            
                mCreatedAt: "created_at",
            
                mUpdatedAt: "updated_at",
            
                mName: "name",
            
                mAge: "age",
            
                mBirthday: "birthday",
            
                mCompanyID: "company_id",
            
                mManagerID: "manager_id",
            
        }
    }

    func (*User) TableName() string {
    return "users"
    }

    // 函数
    
        func (c *GoUser) FieldNameID() string {
            return c.mID
        }
    
        func (c *GoUser) FieldNameUid() string {
            return c.mUid
        }
    
        func (c *GoUser) FieldNameCreatedAt() string {
            return c.mCreatedAt
        }
    
        func (c *GoUser) FieldNameUpdatedAt() string {
            return c.mUpdatedAt
        }
    
        func (c *GoUser) FieldNameName() string {
            return c.mName
        }
    
        func (c *GoUser) FieldNameAge() string {
            return c.mAge
        }
    
        func (c *GoUser) FieldNameBirthday() string {
            return c.mBirthday
        }
    
        func (c *GoUser) FieldNameCompanyID() string {
            return c.mCompanyID
        }
    
        func (c *GoUser) FieldNameManagerID() string {
            return c.mManagerID
        }
    

    // 字段
    
        func (c *GoUser) AddColID() *GoUser {
            c.cols = append(c.cols, c.mID)
            return c
        }
    
        func (c *GoUser) AddColUid() *GoUser {
            c.cols = append(c.cols, c.mUid)
            return c
        }
    
        func (c *GoUser) AddColCreatedAt() *GoUser {
            c.cols = append(c.cols, c.mCreatedAt)
            return c
        }
    
        func (c *GoUser) AddColUpdatedAt() *GoUser {
            c.cols = append(c.cols, c.mUpdatedAt)
            return c
        }
    
        func (c *GoUser) AddColName() *GoUser {
            c.cols = append(c.cols, c.mName)
            return c
        }
    
        func (c *GoUser) AddColAge() *GoUser {
            c.cols = append(c.cols, c.mAge)
            return c
        }
    
        func (c *GoUser) AddColBirthday() *GoUser {
            c.cols = append(c.cols, c.mBirthday)
            return c
        }
    
        func (c *GoUser) AddColCompanyID() *GoUser {
            c.cols = append(c.cols, c.mCompanyID)
            return c
        }
    
        func (c *GoUser) AddColManagerID() *GoUser {
            c.cols = append(c.cols, c.mManagerID)
            return c
        }
    

    func (c *GoUser) AllCols() []string {
        return c.cols
    }

    func (c *GoUser) ResetCols() {
        c.cols = []string{}
    }

    func (c *GoUser) AddAllCols() {
        c.cols = append(c.cols, "*")
    }


type tableColumn struct {
}

func NewTableColumn() *tableColumn {
    return &tableColumn{}
}


    func (*tableColumn) Book() *GoBook {
        return newGoBook()
    }


    func (*tableColumn) User() *GoUser {
        return newGoUser()
    }


