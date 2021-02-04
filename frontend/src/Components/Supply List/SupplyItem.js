import React, {useContext, useState} from 'react';
import './SupplyItem.scss'
import DisplayError from "../Error/DisplayError";
import {userSession} from "../../UserSession";
import {canDelete, canEdit} from "../Permissions/Permissions";

function SupplyItem(props) {
    const [supplyName, setSupplyName] = useState(props.supply.supply)
    const [supplyDesc, setSupplyDesc] = useState(props.supply.desc)
    const [category, setCategory] = useState(props.supply.item_category.String)

    const [editing, setEditing] = useState(false)
    const [savingChanges, setSavingChanges] = useState(false)
    const [deleted, setDeleted] = useState(false)

    const [user] = useContext(userSession)
    const [error, setError] = useState(null)

    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    function submitChanges() {
        setSavingChanges(true)
        let body = {
            "supply": supplyName,
            "desc": supplyDesc,
            "item_category": {String: "", Valid: false},
        }
        if (category !== "") {
            body["item_category"] = {String: category, Valid: true}
        }

        fetch("/api/v1/supply/" + props.supply.id, {
            method: "POST",
            body: JSON.stringify(body)
        })
            .then((resp) => handleErrors(resp, "Unable to create list. Try again later."))
            .then((resp) => {
                props.saveChanges(resp)
                setSavingChanges(false)
            })
            .catch(error => setError(error.toString()));
    }

    function deleteItem() {
        setSavingChanges(true)

        fetch("/api/v1/supply/" + props.supply.id, {
            method: "DELETE",
        })
            .then((resp) => handleErrors(resp, "Unable to create list. Try again later."))
            .then(() => {
                setDeleted(true)
                setSavingChanges(false)
            })
            .catch(error => setError(error.toString()));
    }

    return (
        !deleted &&
        <div className="supply-item">
            {!editing ? (
                <React.Fragment>
                    <div className="supply-name">{supplyName}</div>
                    <div className="supply-desc">{supplyDesc}</div>
                    <div className="category">{category}</div>
                    {canEdit("supply", user) &&
                    <button onClick={() => setEditing(!editing)}>Edit</button>
                    }
                </React.Fragment>
            ) : (
                <div className="editing-supply">
                    <label>
                        Supply
                        <input value={supplyName}
                               placeholder={"Supply"}
                               onChange={(e) => setSupplyName(e.target.value)}/>
                    </label>
                    <label>
                        Description
                        <input value={supplyDesc}
                               placeholder={"Description"}
                               onChange={(e) => setSupplyDesc(e.target.value)}/>
                    </label>
                    <label>
                        Category
                        <input value={category}
                               placeholder={"Category"}
                               onChange={(e) => setCategory(e.target.value)}/>
                    </label>
                    {canEdit("supply", user) &&
                    <button onClick={() => setEditing(!editing)}>Stop Editing</button>
                    }
                </div>
            )}


            {savingChanges ? (
                <React.Fragment>
                    <button disabled={true}>Waiting for server...</button>
                </React.Fragment>
            ) : (
                <React.Fragment>
                    {canEdit("supply", user) && (
                        <button onClick={submitChanges}>
                            Save
                        </button>
                    )}
                    {canDelete("supply", user) && (
                        <button onClick={deleteItem}>
                            Remove
                        </button>
                    )}
                </React.Fragment>

            )}

            {error != null && <DisplayError msg={error}/>}
        </div>
    );

}

export default SupplyItem;
