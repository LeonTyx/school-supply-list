import './Home.scss'
import {useContext, useEffect, useState} from "react";
import SchoolCard from "./SchoolCard";
import CreateSchool from "./CreateSchool";
import {userSession} from "../../UserSession";
import {canCreate} from "../Permissions/Permissions";

function Home() {
    const [schools, setSchools] = useState(null);
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
        fetch("/api/v1/schools")
            .then((resp) => handleErrors(resp, "Unable to retrieve roles"))
            .then(
                (result) => {
                    setSchools(result);
                }, (error) => {
                    setSchools(null)
                    setError(error);
                }
            )
    }, [])

    function addSchool(newSchool){
        let schoolsCopy = Object.assign([], schools)
        schoolsCopy.push(newSchool)

        setSchools(schoolsCopy)
    }

    function removeFromList(id){
        let schoolsCopy = Object.assign([], schools)
        schoolsCopy.forEach((school, index)=>{
            if(school.school_id === id){
                schoolsCopy.splice(index, 1)
            }
        })

        setSchools(schoolsCopy)
    }

    return (
        error === null &&
        <div className="home">
            <section className="schools">
                <h2>Schools</h2>
                {schools !== null && schools.map((school) =>
                    <SchoolCard removeFromList={removeFromList}
                                key={school.school_id}
                                school={school}/>
                )}
            </section>

            {canCreate("school", user) && <CreateSchool addSchool={addSchool}/>}
        </div>
    );
}

export default Home;
