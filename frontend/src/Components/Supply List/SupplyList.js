import './SupplyList.scss'
import SupplyItem from "./SupplyItem";
import {useEffect, useState, useContext} from "react";
import AddSupply from "./AddSupply";
import DisplayError from "../Error/DisplayError";
import {userSession} from "../../UserSession";

function SupplyList(props) {
    const [list, setList] = useState(null)
    const [checked, setChecked] = useState(null)
    const [user] = useContext(userSession)

    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    function saveItem(supply) {
        let updatedList = deleteOldItem(supply.id)
        if (!supply.item_category.Valid) {
            updatedList["basic_supplies"].push(supply)
        } else {
            let targetCategory = supply.item_category.String
            updatedList["categorized_supplies"][targetCategory].push(supply)
        }

        setList(updatedList)
    }
    function createItem(supply){
        let updatedList = JSON.parse(JSON.stringify(list))

        if (!supply.item_category.Valid) {
            updatedList["basic_supplies"].push(supply)
        } else {
            let targetCategory = supply.item_category.String
            if(updatedList["categorized_supplies"][targetCategory] === undefined){
                updatedList["categorized_supplies"][targetCategory] = [];
            }
            updatedList["categorized_supplies"][targetCategory].push(supply)
        }

        setList(updatedList)
    }
    function deleteOldItem(id) {
        let updatedList = JSON.parse(JSON.stringify(list))

        if (updatedList["basic_supplies"] !== null) {
            updatedList["basic_supplies"].forEach((savedSupply, i) => {
                if (savedSupply.id === id) {
                    updatedList["basic_supplies"].splice(i, 1)
                    return updatedList
                }
            })
        }

        if (updatedList["categorized_supplies"] !== null) {
            let categories = Object.keys(updatedList["categorized_supplies"])
            categories.forEach((category) => {
                //Go through each item in each category
                if (updatedList["categorized_supplies"][category] !== null) {
                    updatedList["categorized_supplies"][category].forEach((savedSupply, i) => {
                        if (savedSupply.id === id) {
                            updatedList["categorized_supplies"][category].splice(i, 1)
                            return updatedList
                        }
                    })
                }
            })
        }
        return updatedList
    }

    function toggleCheckbox(e, supplyID){
        let copyOfChecked = Object.assign([], checked)

        if(copyOfChecked.includes(supplyID)){
            copyOfChecked.splice(copyOfChecked.indexOf(supplyID), 1)
        }else{
            copyOfChecked.push(supplyID)
        }
        setChecked(copyOfChecked)
    }

    function updateChecked(){
        fetch("/api/v1/saved-list/" + list.list_id, {
            method: "POST",
            body: JSON.stringify(checked)
        })
            .then((resp) => handleErrors(resp, "Unable to create list. Try again later."))
            .then((resp) => {
                setChecked(resp)
            })
            .catch(error => setError(error.toString()));
    }

    useEffect(() => {
        //Fetch supply list from api
        fetch("/api/v1/supply-list/" + props.match.params.id)
            .then((resp) => handleErrors(resp, "Unable to retrieve this supply list"))
            .then(
                (result) => {
                    setList(result)
                    if(result.checked == null && user != null){
                        setChecked([])
                    }else{
                        setChecked(result.checked)
                    }
                }, (error) => {
                    setList(null)
                    setChecked(null)
                    setError(error.toString())
                }
            )
    }, [props.match.params.id, user])

    return (
        error === null && list !== null ? (
            <div className="supply-list">
                <h2>{list.list_name}</h2>
                <div className="basic-supplies">
                    {checked != null ?
                        list["basic_supplies"] != null &&
                        list["basic_supplies"].map((supply) =>
                            <label key={supply.id}>
                                <input type="checkbox"
                                       onChange={event => toggleCheckbox(event, supply.id)}
                                       checked={checked.includes(supply.id)}/>
                                <SupplyItem saveChanges={saveItem}
                                            supply={supply}/>
                            </label>
                        ):
                        list["basic_supplies"] != null &&
                        list["basic_supplies"].map((supply) =>
                            <SupplyItem saveChanges={saveItem}
                                        key={supply.id}
                                        supply={supply}/>
                        )
                    }
                </div>
                {Object.keys(list["categorized_supplies"]).length > 0 &&
                <div className="categorized">
                    {Object.keys(list["categorized_supplies"]).map((category) =>
                        <div key={category}>
                            <h3>{category}</h3>
                            {checked != null ?
                                list["categorized_supplies"][category] !== null && (
                                    list["categorized_supplies"][category].map((supply) =>
                                        <label key={supply.id}>
                                            <input type="checkbox"
                                                   onChange={event => toggleCheckbox(event, supply.id)}
                                                   checked={checked.includes(supply.id)}/>
                                            <SupplyItem saveChanges={saveItem}
                                                        supply={supply}/>
                                        </label>
                                    )) :
                                list["categorized_supplies"][category] !== null && (
                                    list["categorized_supplies"][category].map((supply) =>
                                        <SupplyItem saveChanges={saveItem}
                                                    key={supply.id}
                                                    supply={supply}/>
                                    ))}
                        </div>
                    )}
                </div>
                }
                <AddSupply listID={list.list_id} addSupply={createItem}/>
                {checked != null && <button onClick={()=>updateChecked()}>Save</button>}
            </div>
        ) : (
            <DisplayError msg={error}/>
        )
    );

}

export default SupplyList;
