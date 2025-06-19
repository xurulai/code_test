// app.go
package sqlmock_demo

import "database/sql"

// recordStats 记录用户浏览产品信息，使用事务确保数据一致性
// 参数：
//   db - 数据库连接
//   userID - 用户ID
//   productID - 产品ID
// 返回：
//   错误信息，成功时为nil
func recordStats(db *sql.DB, userID, productID int64) (err error) {
	// 开启数据库事务
	// 事务用于确保多个SQL操作要么全部成功，要么全部失败
	tx, err := db.Begin()
	if err != nil {
		return // 开启事务失败，直接返回错误
	}

	// 使用defer确保事务在函数结束时正确提交或回滚
	defer func() {
		switch err {
		case nil:
			err = tx.Commit() // 无错误时提交事务
		default:
			tx.Rollback() // 有错误时回滚事务
		}
	}()

	// 更新products表的浏览量
	// 增加对应产品的总浏览次数
	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return // 更新失败，函数返回时defer会回滚事务
	}

	// 在product_viewers表中记录用户浏览行为
	// 记录哪个用户浏览了哪个产品
	if _, err = tx.Exec(
		"INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)",
		userID, productID); err != nil {
		return // 插入失败，函数返回时defer会回滚事务
	}

	return // 所有操作成功，函数返回时defer会提交事务
}

func main() {
	// 打开数据库连接（注意：测试过程中不需要真正连接）
	// 参数：数据库驱动、连接字符串
	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err) // 连接失败，程序崩溃
	}
	defer db.Close() // 确保main函数结束时关闭数据库连接

	// 模拟用户浏览行为：userID为1的用户浏览了productID为5的产品
	if err = recordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
		panic(err) // 记录失败，程序崩溃
	}
}
