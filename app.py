import os
import logging
from flask import Flask, jsonify

# Set up logging for easier debugging
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# Create Flask application
app = Flask(__name__)
app.secret_key = os.environ.get("SESSION_SECRET")

@app.route('/greeting', methods=['GET'])
def greeting():
    """
    Endpoint to return a hello message.
    
    Returns:
        JSON response with a greeting message and HTTP status code 200
    """
    logger.debug("Greeting endpoint called")
    try:
        return jsonify({
            "message": "Hello! Welcome to our greeting service. Have a wonderful day!",
            "status": "success"
        }), 200
    except Exception as e:
        logger.error(f"Error in greeting endpoint: {str(e)}")
        return jsonify({
            "message": "An error occurred while processing your request",
            "status": "error",
            "error": str(e)
        }), 500

# Error handlers for common HTTP errors
@app.errorhandler(404)
def not_found(error):
    """Handle 404 errors"""
    return jsonify({
        "message": "Endpoint not found",
        "status": "error"
    }), 404

@app.errorhandler(405)
def method_not_allowed(error):
    """Handle 405 errors"""
    return jsonify({
        "message": "Method not allowed for this endpoint",
        "status": "error"
    }), 405

@app.errorhandler(500)
def server_error(error):
    """Handle 500 errors"""
    return jsonify({
        "message": "Internal server error",
        "status": "error"
    }), 500
