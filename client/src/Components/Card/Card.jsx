import './card.scss';
function Card({ title, imgUrl, description, rating }) {
  const ratings = document.querySelectorAll('.rating');
  if (ratings.length > 0) {
    initRatings();
  }

  function initRatings() {
    let ratingActive, ratingValue;

    for (let index = 0; index < ratings.length; index++) {
      const rating = ratings[index];
      initRating(rating);
    }

    function initRating(rating) {
      initRatingVars(rating);

      setRatingActiveWidth();
    }

    function initRatingVars(rating) {
      ratingActive = rating.querySelector('.rating__active');
      ratingValue = rating.querySelector('.rating__value');
    }

    function setRatingActiveWidth(index = ratingValue.innerHTML) {
      const setRatingActiveWidth = index / 0.05;
      ratingActive.style.width = `${setRatingActiveWidth}%`;
    }
  }
  return (
    <div className="card">
      <div className="imgBlock">
        <img className="dishImg" src={imgUrl} alt="dish" />
      </div>

      <div className="share">
        <img src="/img/share.svg" alt="share" />
      </div>

      <div className="titleBlock">
        <h2>{title}</h2>
        <span>Served with french fries + drink</span>
      </div>

      <div className="description">
        <span>{description}</span>
      </div>

      <div className="rate-order">
        <div className="rating">
          <div className="rating__body">
            <div className="rating__active"></div>
            <div className="rating__items">
              <input type="radio" className="rating__item" value="1" name="rating" />
              <input type="radio" className="rating__item" value="2" name="rating" />
              <input type="radio" className="rating__item" value="3" name="rating" />
              <input type="radio" className="rating__item" value="4" name="rating" />
              <input type="radio" className="rating__item" value="5" name="rating" />
            </div>
          </div>
          <div className="rating__value">{rating}</div>
        </div>
        <div className="order-button">
          <span>ORDER</span>
        </div>
      </div>
    </div>
  );
}

export default Card;
