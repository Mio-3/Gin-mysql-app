package db

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "todo-app/models"
)

var DB *gorm.DB

func Init() (*gorm.DB, error) {
    // localhostを使用して接続
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        "admin",      // ユーザー名
        "admin123",   // パスワード
        "localhost",  // ホスト
        "3306",      // ポート
        "todo_db",    // データベース名
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect database: %v", err)
    }

    // マイグレーション
    err = db.AutoMigrate(&models.Todo{})
    if err != nil {
        return nil, fmt.Errorf("failed to migrate database: %v", err)
    }

    DB = db
    return db, nil
}