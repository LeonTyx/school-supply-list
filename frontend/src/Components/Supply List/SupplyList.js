import './SupplyList.scss'
import SupplyItem from "./SupplyItem";
import {useEffect, useState} from "react";
import AddSupply from "./AddSupply";
import DisplayError from "../Error/DisplayError";

function SupplyList(props) {
    const [list, setList] = useState(null)

    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    function saveItem(supply){
        let updatedList = deleteOldItem(supply.id)
        if(!supply.item_category.Valid){
            updatedList["basic_supplies"].push(supply)
        }else{
            let targetCategory = supply.item_category.String
            updatedList["categorized_supplies"][targetCategory].push(supply)
        }

        setList(updatedList)
    }

    function deleteOldItem(id){
        let updatedList = JSON.parse(JSON.stringify(list))

        if(updatedList["basic_supplies"] !== null){
            updatedList["basic_supplies"].forEach((savedSupply, i)=>{
                if(savedSupply.id === id){
                    updatedList["basic_supplies"].splice(i, 1)
                    return updatedList
                }
            })
        }

        if(updatedList["categorized_supplies"] !== null){
            let categories = Object.keys(updatedList["categorized_supplies"])
            categories.forEach((category) => {
                //Go through each item in each category
                if(updatedList["categorized_supplies"][category] !== null){
                    updatedList["categorized_supplies"][category].forEach((savedSupply, i)=>{
                        if(savedSupply.id === id){
                            updatedList["categorized_supplies"][category].splice(i, 1)
                            return updatedList
                        }
                    })
                }
            })
        }
        return updatedList
    }

    useEffect(() => {
        //Fetch supply list from api
        fetch("/api/v1/supply-list/" + props.match.params.id)
            .then((resp) => handleErrors(resp, "Unable to retrieve this supply list"))
            .then(
                (result) => {
                    setList(result);
                }, (error) => {
                    setList(null)
                    setError(error.toString())
                }
            )

    }, [props.match.params.id])

    return (
        error === null && list !== null ? (
        <div className="supply-list">
            <div className="basic-supplies">
                {list["basic_supplies"] != null &&
                list["basic_supplies"].map((supply) =>
                    <SupplyItem saveChanges={saveItem}
                                key={supply.id}
                                supply={supply}/>
                )}
            </div>
            {Object.keys(list["categorized_supplies"]).map((category)=>
                <div key={category}>
                    <h2>{category}</h2>
                    {list["categorized_supplies"][category] !== null && (
                        list["categorized_supplies"][category].map((supply) =>
                            <SupplyItem saveChanges={saveItem}
                                        key={supply.id}
                                        supply={supply}/>
                        ))}
                </div>
            )}
            <AddSupply listID={list.list_id}/>
        </div>
        ) : (
            <DisplayError msg={error}/>
        )
    );

}

export default SupplyList;
