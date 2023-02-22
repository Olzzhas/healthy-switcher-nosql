import "./dish.scss"

import Navbar from "../../Components/Navbar/Navbar"

function Dish(){
    return(
        <div className="dish-wrapper">
            <div className="dish-info-block">
                <div className="dish-img-block">
                    <img src="/img/dishes/dish_1.png" alt="meal" />
                </div>
                <div className="dish-info-properties">
                    <h3>Featured Meal</h3>
                    <p>Choice of: Coke, Fanta, Sprite, Upgrade to large fries, Add whopper patty, Add Tender crisp patty and more...</p>
                    <p>Price: 1990$</p>
                    <p>Rating: 5</p>
                    <button>Order</button>
                </div>
            </div>
        </div>
    )
    
}

export default Dish