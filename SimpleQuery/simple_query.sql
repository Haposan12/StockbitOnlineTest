SELECT u.ID, u.UserName, p.UserName as ParentUserName FROM user u LEFT JOIN user p ON u.Parent = p.ID