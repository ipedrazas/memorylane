package main

import "database/sql"

type task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (t *task) getTask(db *sql.DB) error {
	return db.QueryRow("SELECT name, completed FROM tasks WHERE id=$1",
		t.ID).Scan(&t.Name, &t.Completed)
}

func (t *task) updateTask(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE tasks SET name=$1, completed=$2 WHERE id=$3",
			t.Name, t.Completed, t.ID)

	return err
}

func (t *task) deleteTask(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id=$1", t.ID)

	return err
}

func (t *task) createTask(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO tasks(name, completed) VALUES($1, $2) RETURNING id",
		t.Name, t.Completed).Scan(&t.ID)

	if err != nil {
		return err
	}

	return nil
}

func getTasks(db *sql.DB, start, count int) ([]task, error) {
	rows, err := db.Query(
		"SELECT id, name,  completed FROM tasks LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []task{}

	for rows.Next() {
		var t task
		if err := rows.Scan(&t.ID, &t.Name, &t.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
