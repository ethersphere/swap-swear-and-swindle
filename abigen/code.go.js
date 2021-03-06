const fs = require('fs')

const [,,package,compiled,contract] = process.argv

const output = JSON.parse(fs.readFileSync(compiled))

function makeCodeFile(contract, path=`contracts/${contract}.sol`) {
  const Contract = output.contracts[`${path}:${contract}`]
  const binRuntime = Contract['bin-runtime']

  return `  // Copyright 2019 The Swarm Authors
  // This file is part of the Swarm library.
  //
  // The Swarm library is free software: you can redistribute it and/or modify
  // it under the terms of the GNU Lesser General Public License as published by
  // the Free Software Foundation, either version 3 of the License, or
  // (at your option) any later version.
  //
  // The Swarm library is distributed in the hope that it will be useful,
  // but WITHOUT ANY WARRANTY; without even the implied warranty of
  // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
  // GNU Lesser General Public License for more details.
  //
  // You should have received a copy of the GNU Lesser General Public License
  // along with the Swarm library. If not, see <http://www.gnu.org/licenses/>.
  //
  // Code generated - DO NOT EDIT.
  // This file was autogenerated with 'npm run abigen' from ethersphere/swap-swear-and-swindle and any manual changes will be lost.
  package ${package}

  // ${contract}DeployedCode is the bytecode ${contract} will have after deployment
  const ${contract}DeployedCode = "0x${binRuntime}"`
}

console.log(makeCodeFile(contract))