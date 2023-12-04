import React, { useState } from "react";
import axios from "axios";


const Container = {
  backgroundColor: "rgba(255, 255, 255, 1)",
  display: "flex",
  justifyContent: "center", // Center the paper horizontally
  alignItems: "center", // Center the paper vertically
  height: "100vh",
};

const selamatDatangStyle = {
  textAlign: "center",
  color: "rgba(0, 0, 0, 1)",
  fontFamily: "Poppins",
  fontWeight: "500",
  fontSize: "30px",
  marginTop: "20px", // Adjust the margin as needed
  marginBottom: "20px", // Add bottom margin for spacing
};

const PaperLogin = {
  backgroundColor: "rgba(216, 197, 197, 1)",
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  width: "500px",
  padding: "40px", // Added padding for inner spacing
  borderRadius: "30px",
  boxShadow: "0 4px 8px 0 rgba(0,0,0,0.2)", // Optional: added box shadow for better aesthetics
};

const formStyle = {
  display: "flex",
  flexDirection: "column",
  width: "100%", // Form takes the full width of PaperLogin
};

const inputStyle = {
  margin: "10px 0", // Adds margin to top and bottom of each input
  padding: "10px",
  border: "1px solid",
  borderRadius: "30px",
  width: "calc(100% - 20px)", // Adjust width to account for padding
};

const buttonStyle = {
  backgroundColor: "rgba(120, 125, 243, 1)",
  border: "1px solid rgba(0, 0, 0, 1)",
  borderRadius: "40px",
  width: "68px",
  height: "36px",
  margin: "20px auto", // Center the button
};

function Register() {
  // Logic remains the same...
  const [formdata, setFormdata] = useState({
    nama: "",
    username: "",
    password: "",
  });

  ///event handler
  //membuat event handler ketika tombol di klik

  const handlersubmit = (e) => {
    e.preventDefault();

    //pengiriman data menggunakan Axios
    axios
      .post("http://127.0.0.1:4000/register", formdata)
      .then((respone) => {
        console.log(respone.data);
      })
      .catch((error) => {
        console.error(error);
        console.log("data yang eror yaitu", formdata);
      });
  };

  //handlserchange

  const handlerChage = (e) => {
    setFormdata({ ...formdata, [e.target.name]: e.target.value });
  };

  return (
    <div>
      <div style={Container}>
        <div style={PaperLogin}>
          <span style={{ ...selamatDatangStyle, marginTop: "0" }}>
            Selamat Datang silahkan daftar
          </span>
          <form style={formStyle} onSubmit={handlersubmit}>
            <input
              style={inputStyle}
              type="text"
              name="username"
              value={formdata.username}
              onChange={handlerChage}
              placeholder="Username"
            />
            <input
              style={inputStyle}
              type="text"
              name="nama"
              value={formdata.nama}
              onChange={handlerChage}
              placeholder="Nama"
            />
            <input
              style={inputStyle}
              type="password"
              name="password"
              placeholder="Password"
              onChange={handlerChage}
              value={formdata.password}
            />
            <button style={buttonStyle} type="submit">
              Kirim
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}

export default Register;
