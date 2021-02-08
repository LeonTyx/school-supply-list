import React, {useContext, useState} from 'react';
import "./AddSupply.scss";
import {canCreate} from "../Permissions/Permissions";
import {userSession} from "../../UserSession";

function AddSupply(props) {
    const [supplyName, setSupplyName] = useState("")
    const [supplyDesc, setSupplyDesc] = useState("")
    const [category, setCategory] = useState("")
    const [submitting, setSubmitting] = useState(false)
    const [user] = useContext(userSession)

    const [error, setError] = useState(null)

    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    function addSupply() {
        setSubmitting(true)
        let body = {
            "list_id": props.listID,
            "supply": supplyName,
            "desc": supplyDesc,
            "item_category": {String: "", Valid: false},
        }
        if (category !== "") {
            body["item_category"] = {String: category, Valid: true}
        }

        fetch("/api/v1/supply", {
            method: "PUT",
            body: JSON.stringify(body),
        })
            .then((resp) => handleErrors(resp, "Unable to add supply. Try again later."))
            .then((resp) => {
                setSupplyName("")
                setSupplyDesc("")
                setCategory("")
                props.addSupply(resp)
                setSubmitting(false)
            })
            .catch(error => setError(error.toString()));
    }

    return (
        canCreate("supply", user) &&
        <form className="add-supply"
              onSubmit={(e) => e.preventDefault()}>
            <label>
                Supply
                <input value={supplyName}
                       onChange={(e) => setSupplyName(e.target.value)}/>
            </label>
            <label>
                Description
                <textarea value={supplyDesc}
                          onChange={(e) => setSupplyDesc(e.target.value)}/>
            </label>
            <label>
                Category
                <input value={category}
                       onChange={(e) => setCategory(e.target.value)}/>
            </label>
            <button onClick={addSupply}>Add</button>
        </form>
    );

}

export default AddSupply;
