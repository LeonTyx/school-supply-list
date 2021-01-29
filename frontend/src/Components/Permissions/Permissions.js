function canCreate(resource, user){
    if (user !== undefined && user !== null) {
        let resc = user.consolidated_resources
        if (resc[resource] !== undefined && resc[resource].policy.can_view) {
            return true
        }
    }
    return false
}

function canView(resource, user){
    if (user !== undefined && user !== null) {
        let resc = user.consolidated_resources
        if (resc[resource] !== undefined && resc[resource].policy.can_view) {
            return true
        }
    }
    return false
}

function canEdit(resource, user){
    if (user !== undefined && user !== null) {
        let resc = user.consolidated_resources
        if (resc[resource] !== undefined && resc[resource].policy.can_view) {
            return true
        }
    }
    return false
}

function canDelete(resource, user){
    if (user !== undefined && user !== null) {
        let resc = user.consolidated_resources
        if (resc[resource] !== undefined && resc[resource].policy.can_view) {
            return true
        }
    }
    return false
}

export {canView, canCreate, canDelete, canEdit}