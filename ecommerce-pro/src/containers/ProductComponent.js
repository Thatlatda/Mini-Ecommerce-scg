import React from "react";
import { Link } from "react-router-dom";
import { useSelector } from "react-redux";
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
      margin: theme.spacing(1),
    },
  },
}));

const ProductComponent = () => {
  const classes = useStyles();
  const products = useSelector((state) => state.allProducts.products);
  const renderList = products.map((product) => {
    const { product_id, product_name, image, price, category } = product;
    return (

      <div className="four wide column" key={product_id}>
        <Link to={`/product/${product_id}`}>
          <div className="ui link cards">
            <div className="card">
              <div className="image">
                <img src={image} alt={product_name} />
                
              </div>
              <div className="content">
                <div className="header">{product_name}</div>
                <div className="meta price">$ {price}</div>
                <div className={classes.root}>
      <Button variant="outlined" color="primary">
        Add to cart
      </Button>
      <Button variant="outlined" color="secondary">
        Buy
      </Button>
    
    </div>
                
              </div>
            </div>
          </div>
        </Link>
      </div>
    );
  });
  return <>{renderList}</>;
};

export default ProductComponent;
