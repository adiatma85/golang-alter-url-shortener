package user

const (
	createUser = `
	INSERT INTO user (email, username, password, display_name, created_by)
	    VALUES (:email, :username, :password, :display_name, :created_by)`

	getUser = `
	SELECT
	    id,
		email,
	    username,
	    password,
	    display_name,
	    status,
	    created_at,
	    created_by,
	    updated_at,
	    updated_by
	FROM
	    user`

	updateUser = `
	UPDATE
	    user`

	readUserCount = `
	SELECT
	    COUNT(*)
	FROM
	    user`
)
