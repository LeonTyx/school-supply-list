import React,{useState} from 'react';

function AddSupply() {
    const [supplyName, setSupplyName] = useState("")
    const [supplyDesc, setSupplyDesc] = useState("")

    function addSupply(){

    }

    return (
        <form>
            <label>
                <input value={supplyName} onChange={(e)=>setSupplyName(e.target.value)}/>
            </label>
            <label>
                <textarea value={supplyDesc} onChange={(e)=>setSupplyDesc(e.target.value)}/>
            </label>
            <button onClick={addSupply}>Add</button>
        </form>
    );

}

export default AddSupply;
