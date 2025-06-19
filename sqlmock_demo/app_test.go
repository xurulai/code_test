package sqlmock_demo

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestShouldUpdateStats 测试recordStats函数在SQL执行成功时的行为
func TestShouldUpdateStats(t *testing.T) {
	// 创建一个模拟的数据库连接，无需连接真实数据库
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建模拟数据库连接时出错: %s", err)
	}
	defer db.Close() // 测试结束后关闭连接

	// 设置期望的数据库操作序列
	mock.ExpectBegin() // 期望调用Begin()开启事务

	// 期望执行UPDATE语句，并返回结果：LastInsertId=1, RowsAffected=1
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))

	// 期望执行INSERT语句，带参数(2, 3)，并返回结果
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit() // 期望调用Commit()提交事务

	// 调用被测试函数，传入模拟数据库连接和测试参数
	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("更新统计信息时不应出错: %s", err)
	}

	// 验证所有期望的数据库操作都被正确执行
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("存在未满足的期望: %s", err)
	}
}

// TestShouldRollbackStatUpdatesOnFailure 测试recordStats函数在SQL执行失败时的回滚行为
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	// 创建模拟数据库连接
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建模拟数据库连接时出错: %s", err)
	}
	defer db.Close()

	// 设置期望的数据库操作序列（包含错误情况）
	mock.ExpectBegin() // 期望开启事务

	// 期望UPDATE语句成功执行
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))

	// 期望INSERT语句失败，返回自定义错误
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnError(fmt.Errorf("some error"))

	mock.ExpectRollback() // 期望调用Rollback()回滚事务

	// 执行被测试函数
	if err = recordStats(db, 2, 3); err == nil {
		t.Errorf("应该返回错误，但实际没有")
	}

	// 验证所有期望的操作都被执行
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("存在未满足的期望: %s", err)
	}
}
