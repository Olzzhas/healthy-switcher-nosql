import React from 'react';
import { Route, Routes } from 'react-router-dom';
import Navbar from './Components/Navbar/Navbar';
import Main from './Pages/Main';
import Register from './Pages/Register';
import './index.css';

import axios from 'axios';

function App() {
  const [dishes, setDishes] = React.useState([]);

  React.useEffect(() => {
    async function fetchData() {
      try {
        const [itemsResponse] = await Promise.all([
          axios.get('http://localhost:5000/api/topDishes'),
        ]);

        setDishes(itemsResponse.data);
      } catch (error) {
        console.log(error);
      }
    }
    fetchData();
  }, []);
  console.log(dishes);
  return (
    <div className="wrapper">
      <Navbar />
      <Routes>
        <Route exact path="/" element={<Main dishes={dishes} />}></Route>
        <Route exact path="/register" element={<Register />}></Route>
      </Routes>
    </div>
  );
}

export default App;
