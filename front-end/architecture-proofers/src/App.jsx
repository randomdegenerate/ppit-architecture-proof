import { useEffect, useState } from "react";
import "./app.css";

export default function App(){
    const [items, setItems] = useState([]) 

    useEffect(() => {
        fetch('http://localhost:8080/items')
        .then(response => {
            return response.json()
        }).then(jsonData => {
            setItems(jsonData)
        })
        .catch(error => {
            console.log("Error:" +  error)
        })

    }, [])

    return <> 
            <div className="grid center-text inventory-grid"> 
                {items.map(item => {
                    return <div key={item.name} className="inventory-grid-item row">
                            <div className="col">{item.name}</div> 
                            <div className="col">{item.vendor}</div>
                            <div className="col">{item.barcode}</div>
                            <div className="col">{item.price}</div>
                        </div>
                })}
            </div>
        </>
}
