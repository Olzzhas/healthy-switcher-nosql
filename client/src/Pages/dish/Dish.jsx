import { UserContext } from "../../UserContext"
import "./dish.scss"
import axios from "axios"
import React, {useContext} from "react"


function Dish(){

    // async function makeOrder(event) {
    //     event.preventDefault();
    
    //     axios.post("http://localhost:4000/api/tokens/authentication",{
    //         email: email,
    //         password: password,
    
    //     },)
    //     .then(res=>{
    //         setToken(res.data.authentication_token.token)
          
    //     }).then(()=>{
    //       window.location.assign("http://localhost:3000");
    //     })
    //   }

    const {user, setUser} = useContext(UserContext)
    const user_id = user.id
    
    const authHeader = "Bearer " + localStorage.getItem("accessToken")
    console.log(authHeader);

    let dishid = localStorage.getItem("cardId")

    async function makeOrder(){
        axios.post("http://localhost:4000/api/order",{
            dish_id:dishid,
            user_id:user_id,
        },
        {headers:{
            "Authorization": `Bearer ${localStorage.getItem("accessToken")}`
        }})
    }

    const testDish = {
        title:"TestTitle",
        description: "testDescription",
        price: 500,
        rating: 2
    }
    return(
        <div className="dish-wrapper">
            <div className="dish-info-block">
                <div className="dish-img-block">
                    <img src={localStorage.getItem("cardImg")} alt="meal" />
                </div>
                <div className="dish-info-properties">
                    <h3>{localStorage.getItem("cardTitle")}</h3>
                    <p>{localStorage.getItem("cardDescription")}</p>
                    <p>Price: {localStorage.getItem("cardPrice")}$</p>
                    <p>Rating: {localStorage.getItem("cardRating")}</p>
                    <button onClick={()=>{makeOrder()}}>Order</button>
                </div>
            </div>
        </div>
    )
    
}

export default Dish