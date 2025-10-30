package db

import ( "database/sql"
_ "github.com/mattn/go-sqlite3"
"os"
    "log")



func InitDB(dsn string) (*sql.DB, error) {
    // Abrir la conexión
    db, err := sql.Open("sqlite3", dsn)
    if err != nil {
        return nil, err
    }

    // Probar que realmente se pueda conectar
    if err := db.Ping(); err != nil {
        return nil, err
    }

    // Si todo sale bien, retornamos la conexión
    return db, nil
}



func RunMigrations(db *sql.DB) error{
	entries, err := os.ReadDir("internal/db/migrations")
	if err != nil {
        return err
    }


	for _,ent := range entries {
		name := ent.Name()
	    sqlBytes, err := os.ReadFile("internal/db/migrations/" + ent.Name())
        if err != nil {
            return err
        }

        _, err = db.Exec(string(sqlBytes))
        if err != nil {
            return err
        }
        log.Printf("✅ Migración ejecutada: %s", name)
}

return nil
}
