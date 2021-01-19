DELETE
FROM role_resource_bridge
WHERE role_id = 2;
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, true, true, true, true, 1, 2);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, true, true, true, true, 2, 2);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, true, true, true, true, 3, 2);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, true, true, true, true, 4, 2);
INSERT INTO role_resource_bridge (rrb_id, can_add, can_view, can_edit, can_delete, resource_id, role_id)
VALUES (default, true, true, true, true, 5, 2);