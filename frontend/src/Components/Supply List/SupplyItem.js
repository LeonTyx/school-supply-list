import React from 'react';
import './SupplyItem.scss'

function SupplyItem(props) {
    return (
        <div className="supply-item">
            <div className="supply-name">{props.list.supply}</div>
            <div className="supply-desc">{props.list.supply}</div>
            <button>
                Remove
            </button>
        </div>
    );

}

export default SupplyItem;
