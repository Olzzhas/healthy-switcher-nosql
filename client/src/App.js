import React from 'react';
import { json, Route, Routes } from 'react-router-dom';

import Navbar from './Components/Navbar/Navbar';
import Main from './Pages/main/Main';
import Register from './Pages/register/Register';
import Login from './Pages/login/register/Login';
import Dish from './Pages/dish/Dish';



import './index.css';

import axios from 'axios';
import { UserContext } from './UserContext';

function App() {
  const [dishes, setDishes] = React.useState([]);
  const [token, setToken] = React.useState(localStorage.getItem('accessToken'))

  // let checkAuth
  // if(localStorage.getItem("currentUser") === ""){
  //   checkAuth = false
  //   console.log("user authed");
  // }else{
  //   checkAuth = true
  //   console.log("user unauthed");
  // }

  // const [authorized, setAuthorized] = React.useState(checkAuth)

  React.useEffect(() => {
    async function fetchUser(){
      try {
        localStorage.setItem('accessToken', token);
        if(token === ""){
          return
        }else{
          const [userResponse] = await Promise.all([
            axios.post('http://localhost:4000/api/user/byToken',{
              token:localStorage.getItem("accessToken")
            })
          ])
          
          // console.log(userResponse.data.user);
          // localStorage.setItem("currentUser", JSON.stringify(userResponse.data.user))
          // let test = JSON.parse(localStorage.getItem("currentUser"))
          // console.log(test);
        }
      } catch (error) {
        throw error
      }
    }
    fetchUser()

  }, [token]);


  React.useEffect(() => {
    async function fetchData() {
      try {
        const [itemsResponse] = await Promise.all([
          axios.get('http://localhost:4000/api/topDishes'),
        ]);

        setDishes(itemsResponse.data.dishes);
      } catch (error) {
        console.log(error);
      }
    }
    fetchData();
  }, []);
  console.log(dishes);

  

  function testToken(){
    console.log(localStorage.getItem("accessToken"));
    console.log(localStorage.getItem("currentUser"));
    // console.log(authorized);
  }


  return (
    <UserContext.Provider value={JSON.parse(localStorage.getItem("currentUser"))}>
      <div className="wrapper">
        <Navbar token={token} />
        <button onClick={()=>testToken()}>test token</button>
        <Routes>
          <Route exact path="/" element={<Main dishes={dishes} />}></Route>
          <Route exact path="/register" element={<Register />}></Route>
          <Route exact path="/login" element={<Login token={token} setToken={setToken} />}></Route>
          <Route exact path='/dish' element={<Dish token={token}/>}></Route>
        </Routes>
        
      </div>
    </UserContext.Provider>
  );
}

export default App;
