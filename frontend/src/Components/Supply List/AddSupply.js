import React,{useState, useContext} from 'react';
import {canCreate} from "../Permissions/Permissions";
import {userSession} from "../../UserSession";

function AddSupply() {
    const [supplyName, setSupplyName] = useState("")
    const [supplyDesc, setSupplyDesc] = useState("")
    const [submitting, setSubmitting] = useState(false)
    const [user] = useContext(userSession)

    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response;
    }
    function addSupply() {
        setSubmitting(true)
        fetch("/api/v1/supply", {
            method: "PUT",
            body: JSON.stringify({
                "supply": supplyName,
                "desc": supplyDesc,
            })
        })
            .then((resp) => handleErrors(resp, "Unable to add supply. Try again later."))
            .then(response => {
                setSupplyName("")
                setSupplyDesc("")
                setSubmitting(false)
            })
            .catch(error => setError(error.toString()));
    }

    return (
        canCreate("supply", user) &&
        <form>
            <label>
                <input value={supplyName} onChange={(e)=>setSupplyName(e.target.value)}/>
            </label>
            <label>
                <textarea value={supplyDesc} onChange={(e)=>setSupplyDesc(e.target.value)}/>
            </label>
            
            <button onClick={addSupply}>Add</button>
        </form>
    );

}

export default AddSupply;
