SELECT 
    u.id,
    u.email,
    u.phone,
    u.type,
    u.created_at,
    u.updated_at,
    u.deleted_at
FROM users u
WHERE u.deleted_at is null
