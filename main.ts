/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-13
 * @fileoverview program to make calculator 
 */

// set variables 
//let userOperation: string; 
let answer: number;

// get input for operation 
const userOperation = prompt("Welcome to my calculator program.\nwhich operation would you like to perform today? (select by typing the letter in the front of the operation.)\na.add\nb.subtract\nc.multiply\nd.divide\ne.absolute value\nf.round\ng.raise to an exponent\nh.square root\n") || "";

// if statements for each operation 
if (userOperation == "a"){
  let userValue1AsNumber: number;
  let userValue2AsNumber: number;
  const userValue1AsString = prompt("Please enter your first number") || "0";
  const userValue2AsString = prompt("Please enter your second number") || "0";
  userValue1AsNumber = parseInt(userValue1AsString);
  userValue2AsNumber = parseInt(userValue2AsString);
  answer = userValue1AsNumber + userValue2AsNumber
  console.log("The sum of ",userValue1AsNumber," and ",userValue2AsNumber," is ",answer);
}

if (userOperation == "b") {
  let userValue1AsNumber: number;
  let userValue2AsNumber: number;
  const userValue1AsString = prompt("Please enter your first number") || "0";
  const userValue2AsString = prompt("Please enter your second number") || "0";
  userValue1AsNumber = parseInt(userValue1AsString);
  userValue2AsNumber = parseInt(userValue2AsString);
  answer = userValue1AsNumber - userValue2AsNumber
  console.log("The difference between ", userValue1AsNumber, " and ", userValue2AsNumber, " is ", answer);
}

if (userOperation == "c") {
  let userValue1AsNumber: number;
  let userValue2AsNumber: number;
  const userValue1AsString = prompt("Please enter your first number") || "0";
  const userValue2AsString = prompt("Please enter your second number") || "0";
  userValue1AsNumber = parseInt(userValue1AsString);
  userValue2AsNumber = parseInt(userValue2AsString);
  answer = userValue1AsNumber * userValue2AsNumber
  console.log("The product of ", userValue1AsNumber, " and ", userValue2AsNumber, " is ", answer);
}

if (userOperation == "d") {
  let userValue1AsNumber: number;
  let userValue2AsNumber: number;
  const userValue1AsString = prompt("Please enter your first number") || "0";
  const userValue2AsString = prompt("Please enter your second number") || "0";
  userValue1AsNumber = parseInt(userValue1AsString);
  userValue2AsNumber = parseInt(userValue2AsString);
  answer = userValue1AsNumber / userValue2AsNumber
  console.log("The quotient of ", userValue1AsNumber, " and ", userValue2AsNumber, " is ", answer);
}

if (userOperation == "e") {
  let userValue1AsNumber: number;
  const userValue1AsString = prompt("Please enter the integer to make absolute") || "0";
  userValue1AsNumber = parseInt(userValue1AsString);
  answer = Math.abs(userValue1AsNumber);
  console.log("The absolute value of ", userValue1AsNumber," is ", answer);
}

if (userOperation == "f") {
  let userValue1AsNumber: number;
  const userValue1AsString = prompt("Please enter the number you want to round") || "0";
  userValue1AsNumber = parseFloat(userValue1AsString);
  answer = Math.round(userValue1AsNumber);
  console.log("The rounded value of ", userValue1AsNumber, " is ", answer);
}

if (userOperation == "g") {
  let userValue1AsNumber: number;
  let userValue2AsNumber: number;
  const userValue1AsString = prompt("Please enter your base number") || "0";
  const userValue2AsString = prompt("Please enter your exponent") || "0";
  userValue1AsNumber = parseInt(userValue1AsString);
  userValue2AsNumber = parseInt(userValue2AsString);
  answer = Math.pow(userValue1AsNumber, userValue2AsNumber)
  console.log("The power of ", userValue1AsNumber, " to ", userValue2AsNumber, " is ", answer);
}

if (userOperation == "h") {
  let userValue1AsNumber: number;
  const userValue1AsString = prompt("Please enter the number you want to round") || "0";
  userValue1AsNumber = parseFloat(userValue1AsString);
  answer = Math.sqrt(userValue1AsNumber);
  console.log("The square root of ", userValue1AsNumber, " is ", answer);
}
console.log("\nDone.");