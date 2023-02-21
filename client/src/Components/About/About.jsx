import './about.scss';
function About() {
  return (
    <div className="about-wrapper">
      <div>
        <h1>ABOUT</h1>
      </div>
      <span className="spanAbout">The Basics Of Healthy Food</span>
      <p align="center" className="lorem">
        In aliqua ea ullamco ad est ex non deserunt nulla. Consectetur sint ea aliquip aliquip
        consectetur voluptate est. Eu minim dolore laboris enim mollit voluptate irure esse aliquip.
      </p>

      <div className="slider">
        <ul>
          <li className="arrowButton">
            <img src="/img/button_left.png" alt="button" />
          </li>
          <li className="dish">
            <img src="/img/dishes/dish_1.png" alt="dish" />
          </li>
          <li className="dish">
            <img src="/img/dishes/dish_2.png" alt="dish" />
          </li>
          <li className="arrowButton">
            <img src="/img/button_right.png" alt="button" />
          </li>
        </ul>
      </div>
    </div>
  );
}

export default About;
