import React, {useEffect, useState} from 'react';
import './School.scss'
import DisplayError from "../Error/DisplayError";
import CreateList from "./CreateList";
import {Link} from "react-router-dom";

function School(props) {
    const [school, setSchool] = useState(null)

    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    useEffect(() => {
        //Fetch school from api
        fetch("/api/v1/school/" + props.match.params.id)
            .then((resp) => handleErrors(resp, "Unable to retrieve school information"))
            .then(
                (result) => {
                    setSchool(result);
                }, (error) => {
                    setSchool(null)
                    setError(error);
                }
            )

    }, [props.match.params.id])

    function removeList(id){
        fetch("/api/v1/supply-list/"+id, {
            method: "DELETE"
        })
            .then((resp) => handleErrors(resp, "Unable to delete supply list"))
            .then(() => {
                let schoolCopy = Object.assign({}, school)
                let supplyListsCopy = Object.assign([], school.supply_lists)
                supplyListsCopy.forEach((list, index) => {
                    if(id === list.list_id){
                        supplyListsCopy.splice(index, 1)
                    }
                })
                schoolCopy["supply_lists"] = supplyListsCopy

                setSchool(schoolCopy)
            })
            .catch(error => setError(error.toString()));
    }

    function addList(list){
        let schoolCopy = Object.assign({}, school)
        let supplyListsCopy = Object.assign([], school.supply_lists)
        supplyListsCopy.push(list)
        schoolCopy["supply_lists"] = supplyListsCopy
        setSchool(schoolCopy)
    }

    return (
        error == null ? (
            school != null && <div className="school">
                <div className="school-header">
                    <h2>{school.school_name}</h2>
                </div>
                <div className="supply-lists">
                    {school.supply_lists != null ? (
                        school.supply_lists.map((list) =>
                            <div className="supply-list-card" key={list.list_id}>
                                <Link to={"/supply-list/" + list.list_id}>
                                    {list.list_name}
                                </Link>
                                <button className="delete"
                                        onClick={()=>removeList(list.list_id)}>
                                    Remove
                                </button>
                            </div>
                        )
                    ) : (
                        <div>No lists yet!</div>
                    )}

                    <CreateList addList={addList} schoolID={parseInt(props.match.params.id)}/>
                </div>
            </div>
        ) : (
            <DisplayError msg={error}/>
        )
    );

}

export default School;
