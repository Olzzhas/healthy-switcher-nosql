import About from '../../Components/About/About';
import Card from '../../Components/Card/Card';
import './main.scss';
function Main({ dishes }) {
  return (
    <div>
      <div className="title">
        <div className="firstLine">
          <span className="light">Your </span>
          <span className="bold">Favorite Food</span>
        </div>
        <br />
        <div className="secondLine">
          <span className="light">Delivered </span>
          <span className="bold">Hot & Fresh</span>
        </div>
        <div className="thirdLine">
          <span>
            HEALTHY SWITCHER chefs do all the prep work, like peeling, chopping & marinating, so you
            can cook a fresh homemade dinner in just 15 minutes.
          </span>
        </div>
      </div>
      <div className="button">
        <p>Order Now</p>
      </div>
      <div>
        <ul className="info">
          <li className="info-element">
            <img src="/img/watch.png" alt="watch" />
            <br />
            <span>Today 10:00 am - 7:00 pm</span>
            <br />
            <span className="text">Working hours</span>
          </li>

          <li className="info-element">
            <img src="/img/location.png" alt="location" />
            <br />
            <span>Alikhan Bokeikhan 25</span>
            <br />
            <span className="text">Get Directions</span>
          </li>

          <li className="info-element">
            <img src="/img/phone.png" alt="phone" />
            <br />
            <span>+7 (747) 574 64 02</span>
            <br />
            <span className="text">Call Online</span>
          </li>
        </ul>
      </div>
      <About />
      <div className="instruction">
        <div className="work">
          <h1>WORK</h1>
          <span className="spanAbout">How It Works</span>
          <img src="/img/green-work-line.png" alt="line" />
        </div>
      </div>
      <div className="qqq"></div>
      <div className="howWorks">
        <ul>
          <li>
            <div className="panel">
              <div className="panelSpan">
                <span>Pick Meals</span>
                <br />
                <img src="/img/panel-line.svg" alt="line" />
              </div>
              <p>
                Choose your meals from our diverse weekly menu. Find gluten or dairy free, low carb
                & veggie options.
              </p>
            </div>
          </li>
          <li>
            <div className="panel">
              <div className="panelSpan">
                <span>Choose How Often</span>
                <br />
                <img src="/img/panel-line.svg" alt="line" />
              </div>
              <p>
                Our team of chefs do the prep work - no more chopping, measuring, or sink full of
                dishes!
              </p>
            </div>
          </li>
          <li>
            <div className="panel">
              <div className="panelSpan">
                <span>Fast Deliveries</span>
                <br />
                <img src="/img/panel-line.svg" alt="line" />
              </div>
              <p>
                Your freshly prepped 15-min dinner kits arrive on your doorstep in a refrigerated
                box.
              </p>
            </div>
          </li>
          <li>
            <div className="panel">
              <div className="panelSpan">
                <span>Tasty Meals</span>
                <br />
                <img src="/img/panel-line.svg" alt="line" />
              </div>
              <p>Gobble makes cooking fast, so you have more time to unwind and be with family.</p>
            </div>
          </li>
        </ul>
      </div>
      <div className="work">
        <h1>DISHES</h1>
        <span className="spanAbout">Dish Of The Day</span>
        <img src="/img/green-work-line.png" alt="line" />
      </div>
      <div className="qqqq"></div>
      <div className="dishBlock">
        {dishes.map((dish) => (
          <Card
            id={dish.id}
            title={dish.title}
            imgUrl={dish.img}
            description={dish.description}
            rating={dish.rating}
          />
        ))}
      </div>
      <div className="work">
        <h1>CHEFS</h1>
        <span className="spanAbout">This Month's Chefs</span>
        <img src="/img/green-work-line.png" alt="line" />
      </div>
      <div className="qqqq"></div>
      <div className="chefBlock">
        <div className="topChefs">
          <div className="chef">
            <div className="chefImg">
              <img src="/img/chefs/gregory.png" alt="chef" />
            </div>
            <div className="chefName">
              <span>Gregory Flores</span>
              <h2>Chef of the cold</h2>
            </div>
          </div>

          <div className="chef">
            <div className="chefImg">
              <img src="/img/chefs/anette.png" alt="chef" />
            </div>
            <div className="chefName">
              <span>Annette Cooper</span>
              <h2>Chef of the hot</h2>
            </div>
          </div>

          <div className="chef">
            <div className="chefImg">
              <img src="/img/chefs/greg.png" alt="chef" />
            </div>
            <div className="chefName">
              <span>Greg Fox</span>
              <h2>Сhef macro kitchen</h2>
            </div>
          </div>
        </div>

        <div className="mealBlock">
          <img src="/img/meal-pic/meal_1.png" alt="meal" />
          <img src="/img/meal-pic/meal_2.png" alt="meal" />
          <img src="/img/meal-pic/meal_3.png" alt="meal" />
          <img src="/img/meal-pic/meal_4.png" alt="meal" />
          <img src="/img/meal-pic/meal_5.png" alt="meal" />
          <img src="/img/meal-pic/meal_6.png" alt="meal" />
          <img src="/img/meal-pic/meal_7.png" alt="meal" />
          <img src="/img/meal-pic/meal_8.png" alt="meal" />
          <img src="/img/meal-pic/meal_9.png" alt="meal" />
        </div>
      </div>

      <div className="work">
        <h1>RECIPES</h1>
        <span className="spanAbout">Recipes From Our Chefs</span>
        <img src="/img/green-work-line.png" alt="line" />
      </div>

      <div className="qqqq"></div>

      <div className="recipes">
        <div className="leftPic">
          {/* <img src="/img/meal-pic/meal_10.png" alt="meal" /> */}
          <div className="leftPic-content">
            <div className="breakfast">
              <span>BREAKFAST</span>
            </div>
            <div>
              <h3>05 Jul 2016</h3>
            </div>
            <div>
              <h2>Lorem ipsum dolor sit amet, consectetur adipiscing elit</h2>
            </div>
            <div>
              <h3>Jason Keller</h3>
            </div>

            <ul className="view-comment">
              <li>
                <img src="/img/views.svg" alt="views" />
                <span>231</span>
              </li>
              <li>
                <img src="/img/comments.svg" alt="views" />
                <span>30</span>
              </li>
            </ul>
          </div>
        </div>

        <ul className="rightSide">
          <li>
            <div className="recipeText">
              <span>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
                incididunt ut
              </span>
            </div>
          </li>

          <li>
            <img src="/img/meal-pic/meal_11.png" alt="meal" />
            <div className="recipeText">
              <span>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
                incididunt ut
              </span>
            </div>
          </li>

          <li>
            <div className="recipeText">
              <span>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
                incididunt ut
              </span>
            </div>
          </li>
        </ul>
      </div>
      <footer>
        <ul className="info">
          <li className="info-element">
            <img src="/img/watch.png" alt="watch" />
            <br />
            <span>Today 10:00 am - 7:00 pm</span>
            <br />
            <span className="text">Working hours</span>
          </li>

          <li className="info-element">
            <img src="/img/location.png" alt="location" />
            <br />
            <span>Alikhan Bokeikhan 25</span>
            <br />
            <span className="text">Get Directions</span>
          </li>

          <li className="info-element">
            <img src="/img/phone.png" alt="phone" />
            <br />
            <span>+7 (747) 574 64 02</span>
            <br />
            <span className="text">Call Online</span>
          </li>
        </ul>
        <div className="footerLine"></div>
        <div className="footerEnd">
          <div className="qwer">
            <div className="logoBottom">
              <img src="/img/Logo.png" alt="logo" />
              <img className="hs-text" src="/img/hs-title.png" alt="title" />
            </div>
            <div className="footerSpan">
              <span>© Designed by Yellow Snow. All Rights Reserved. </span>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}

export default Main;
