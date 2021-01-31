import React from 'react';
import './SupplyItem.scss'

function SupplyItem(props) {

    function editItem(){

    }

    function deleteItem(){

    }
    return (
        <div className="supply-item">
            <div className="supply-name">{props.list.supply}</div>
            <div className="supply-desc">{props.list.supply}</div>

            <button onClick={editItem}>
                Edit
            </button>
            <button onClick={deleteItem}>
                Remove
            </button>
        </div>
    );

}

export default SupplyItem;
