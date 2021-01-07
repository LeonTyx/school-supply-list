import './SupplyList.scss'
import SupplyItem from "./SupplyItem";
import {useEffect, useState} from "react";

function SupplyList() {
    const [list, setList] = useState(null);
    const [error, setError] = useState(null)

    useEffect(() => {
        //Fetch user from api
        fetch("/oauth/v1/supply-list")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                }
            })
            .then(
                (result) => {
                    setList(result);
                }, (error) => {
                    setList(null)
                    setError(error);
                }
            )
    }, [])

    return (
        error === null && list !== null &&
        <div className="supply-list">
            {list.map((supply) =>
                <SupplyItem supply={supply}/>
            )}
        </div>
    );

}

export default SupplyList;
