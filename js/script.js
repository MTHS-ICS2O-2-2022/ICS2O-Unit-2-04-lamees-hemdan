// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Lamees Hemdan
// Created on: March 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates area of a triangle
 */
function calculate() {
  // input
  const Base = parseInt(document.getElementById('base').value)
  const Height = parseInt(document.getElementById('height').value)

  // process
  const area = (Base * Height) / 2

  // output
  document.getElementById('area').innerHTML = 'The Area is ' + area + ' cm²'
}