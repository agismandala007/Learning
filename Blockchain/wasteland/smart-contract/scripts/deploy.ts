import { ethers } from "hardhat";

async function main() {
  // Get the contract factory
  const TrashToken = await ethers.getContractFactory("TrashToken");

  // Specify the initial supply (for example, 1000 tokens)
  const initialSupply = ethers.parseUnits("1000", 18); // Adjust decimals as necessary

  const trashToken = await TrashToken.deploy(initialSupply);

  await trashToken.deployed();

  console.log(`TrashToken deployed to: ${trashToken.address}`);
}

// Run the main function and handle errors
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
