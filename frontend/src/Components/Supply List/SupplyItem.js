import React from 'react';
import './SupplyItem.scss'

function SupplyItem(props) {
    return (
        <div>
            {props.list.supply}
        </div>
    );

}

export default SupplyItem;
