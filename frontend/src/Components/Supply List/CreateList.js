import React, {useContext, useState} from 'react';
import './CreateList.scss'
import DisplayError from "../Error/DisplayError";
import {userSession} from "../../UserSession";
import {canCreate} from "../Permissions/Permissions";

function CreateList(props) {
    const [title, setTitle] = useState("")
    const [grade, setGrade] = useState("")
    const [submitting, setSubmitting] = useState(false)
    const [user] = useContext(userSession)

    const [error, setError] = useState(null)

    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    function submitForm() {
        setSubmitting(true)

        fetch("/api/v1/supply-list", {
            method: "PUT",
            body: JSON.stringify({
                "school_id": props.schoolID,
                "list_name": title,
                "grade": parseInt(grade),
            })
        })
            .then((resp) => handleErrors(resp, "Unable to create list. Try again later."))
            .then(response => {
                setTitle("")
                setGrade("")
                setSubmitting(false)
                props.addList(response)
            })
            .catch(error => setError(error.toString()));
    }

    return (
        canCreate("supply-list", user) &&
        <form className="create-list"
              onSubmit={(e) => e.preventDefault()}>
            <h2>Create List</h2>
            <label>
                List Title
                <input value={title}
                       placeholder="list title"
                       onChange={(e) => setTitle(e.target.value)}
                />
            </label>
            <label>
                Grade (Use -1 for preschool)
                <input value={grade}
                       placeholder="Grade level"
                       type="number"
                       min={-1}
                       max={12}
                       onChange={(e) => setGrade(e.target.value)}/>
            </label>
            {submitting ? (
                <button disabled={true}>Submitting...</button>
            ) : (
                <button onClick={submitForm}>Create list</button>
            )}
            {error != null && <DisplayError msg={error}/>}
        </form>
    );

}

export default CreateList;
