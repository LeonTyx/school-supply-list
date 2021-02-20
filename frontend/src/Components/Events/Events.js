import React, {useContext, useEffect, useState} from 'react';
import './Events.scss'
import {userSession} from "../../UserSession";
import DisplayError from "../Error/DisplayError";

function Events() {
    const [events, setEvents] = useState(null);
    const [error, setError] = useState(null)
    const [user] = useContext(userSession)

    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    useEffect(() => {
        //Fetch user from api
        fetch("/api/v1/events")
            .then((resp) => handleErrors(resp, "Unable to retrieve events"))
            .then(
                (result) => {
                    setEvents(result);
                }, (error) => {
                    setEvents(null)
                    setError(error);
                }
            )
    }, [])

    return (
        <div>
            {events.map((event)=>
                <div>{event.title}</div>
            )}
            {error != null && <DisplayError msg={error}/>}
        </div>
    );

}

export default Events;
