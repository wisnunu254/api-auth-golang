UPDATE users SET 
    email = CASE WHEN ? != '' THEN ? ELSE email END,
    phone = CASE WHEN ? != '' THEN ? ELSE phone END,
    type = CASE WHEN ? != '' THEN ? ELSE type END,
    updated_at = ?
WHERE id = ?