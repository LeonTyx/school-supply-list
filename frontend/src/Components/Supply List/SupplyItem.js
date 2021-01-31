import React, {useState} from 'react';
import './SupplyItem.scss'
import DisplayError from "../Error/DisplayError";

function SupplyItem(props) {
    const [supplyName, setSupplyName] = useState(props.supply.supply)
    const [supplyDesc, setSupplyDesc] = useState(props.supply.desc)
    const [editing, setEditing] = useState(false)
    const [deleting, setDeleting] = useState(false)
    const [deleted, setDeleted] = useState(false)

    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response;
    }

    function editItem(){
        setEditing(true)

        fetch("/api/v1/supply/" + props.supply.id, {
            method: "POST",
            body: JSON.stringify({
                "school_id": supplyName,
                "list_name": supplyDesc,
            })
        })
            .then((resp) => handleErrors(resp, "Unable to create list. Try again later."))
            .then(() => {
                setSupplyName("")
                setSupplyDesc("")
                setEditing(false)
            })
            .catch(error => setError(error.toString()));
    }

    function deleteItem(){
        setDeleting(true)

        fetch("/api/v1/supply/" + props.supply.id, {
            method: "DELETE",
        })
            .then((resp) => handleErrors(resp, "Unable to create list. Try again later."))
            .then(() => {
                setDeleted(true)
                setDeleting(false)
            })
            .catch(error => setError(error.toString()));
    }

    return (
        !deleted &&
        <div className="supply-item">
            <div className="supply-name">{props.list.supply}</div>
            <div className="supply-desc">{props.list.supply}</div>

            {editing ? (
                <button disabled={true}>Editing...</button>
            ) : (
                <button onClick={editItem}>
                    Edit
                </button>
            )}

            {deleting ? (
                <button disabled={true}>Deleting...</button>
            ) : (
                <button onClick={deleteItem}>
                    Remove
                </button>
            )}

            {error != null && <DisplayError msg={error}/>}
        </div>
    );

}

export default SupplyItem;
