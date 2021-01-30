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
            {list.map((supply) =>
                <SupplyItem supply={supply}/>
            )}
            <AddSupply/>
        </div>
        ) : (
            <DisplayError msg={error}/>
        )
    );

}

export default SupplyList;
