DELETE
FROM role_resource_bridge
WHERE role_id = 2;
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, false, true, false, false, 1, 1);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, false, true, false, false, 2, 1);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, false, true, false, false, 3, 1);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, false, false, true, false, 4, 1);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, false, true, false, false, 5, 1);