import React from 'react';
import './SupplyItem.scss'

function SupplyItem(props) {
    return (
        <div>
            {props.list.item_name}
        </div>
    );

}

export default SupplyItem;
