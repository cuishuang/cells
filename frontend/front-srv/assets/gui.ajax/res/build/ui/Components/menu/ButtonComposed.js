/*
 * Copyright 2007-2018 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
 * This file is part of Pydio Cells.
 *
 * Pydio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */
'use strict';

exports.__esModule = true;

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var _Utils = require('./Utils');

var _Utils2 = _interopRequireDefault(_Utils);

var _MenuItemsConsumer = require('./MenuItemsConsumer');

var _MenuItemsConsumer2 = _interopRequireDefault(_MenuItemsConsumer);

var _react = require("react");

var _react2 = _interopRequireDefault(_react);

var _reactDom = require("react-dom");

var _reactDom2 = _interopRequireDefault(_reactDom);

var _materialUi = require("material-ui");

var ButtonComposed = (function (_React$Component) {
    _inherits(ButtonComposed, _React$Component);

    function ButtonComposed() {
        var _this = this;

        _classCallCheck(this, ButtonComposed);

        _React$Component.apply(this, arguments);

        this.state = { showMenu: false };

        this.showMenu = function (event) {
            var anchor = undefined;
            if (event) {
                anchor = event.currentTarget;
            } else {
                anchor = _this._buttonDOM;
            }
            _this.setState({
                showMenu: true,
                anchor: anchor
            });
        };

        this.menuClicked = function (event, index, object) {
            _this.setState({ showMenu: false });
        };
    }

    ButtonComposed.prototype.componentDidMount = function componentDidMount() {
        var _this2 = this;

        if (this.props.openOnEvent) {
            this.props.pydio.observe(this.props.openOnEvent, function () {
                _this2.showMenu();
            });
        }
    };

    ButtonComposed.prototype.render = function render() {
        var _this3 = this;

        var masterButton = undefined,
            arrowButton = undefined;
        var _props = this.props;
        var id = _props.id;
        var masterAction = _props.masterAction;
        var buttonTitle = _props.buttonTitle;
        var primary = _props.primary;
        var secondary = _props.secondary;
        var disabled = _props.disabled;
        var raised = _props.raised;
        var menuItems = _props.menuItems;
        var buttonStyle = _props.buttonStyle;
        var buttonLabelStyle = _props.buttonLabelStyle;
        var className = _props.className;
        var direction = _props.direction;

        var masterLabelStyle = _extends({}, buttonLabelStyle);
        var arrowLabelStyle = _extends({}, buttonLabelStyle);
        if (masterLabelStyle.paddingRight) {
            masterLabelStyle.paddingRight /= Math.floor(3);
        } else {
            masterLabelStyle.paddingRight = 8;
        }
        if (arrowLabelStyle.paddingLeft) {
            arrowLabelStyle.paddingLeft /= Math.floor(3);
        } else {
            arrowLabelStyle.paddingLeft = 8;
        }
        arrowLabelStyle.paddingRight = arrowLabelStyle.paddingLeft;

        var masterProps = {
            primary: primary,
            secondary: secondary,
            disabled: disabled,
            label: buttonTitle,
            style: _extends({}, buttonStyle, { minWidth: 60 }),
            labelStyle: masterLabelStyle,
            onTouchTap: masterAction,
            onClick: function onClick(e) {
                return e.stopPropagation();
            }
        };
        var arrowProps = {
            primary: primary,
            secondary: secondary,
            disabled: disabled,
            label: _react2['default'].createElement('span', { className: "mdi mdi-menu-down" }),
            onTouchTap: this.showMenu,
            style: _extends({}, buttonStyle, { minWidth: 16 }),
            labelStyle: arrowLabelStyle,
            onClick: function onClick(e) {
                return e.stopPropagation();
            }
        };
        var _state = this.state;
        var showMenu = _state.showMenu;
        var anchor = _state.anchor;
        var over = _state.over;

        if (menuItems.length) {
            if (raised) {
                arrowButton = _react2['default'].createElement(_materialUi.RaisedButton, _extends({}, arrowProps, { ref: function (b) {
                        _this3._buttonDOM = _reactDom2['default'].findDOMNode(b);
                    } }));
                masterButton = _react2['default'].createElement(_materialUi.RaisedButton, masterProps);
            } else {
                arrowButton = _react2['default'].createElement(_materialUi.FlatButton, _extends({}, arrowProps, { ref: function (b) {
                        _this3._buttonDOM = _reactDom2['default'].findDOMNode(b);
                    } }));
                masterButton = _react2['default'].createElement(_materialUi.FlatButton, masterProps);
            }
        }
        return _react2['default'].createElement(
            'span',
            { id: id, className: className,
                onMouseOver: function () {
                    _this3.setState({ over: true });
                }, onMouseOut: function () {
                    _this3.setState({ over: false });
                },
                style: over || showMenu ? { backgroundColor: 'rgba(153, 153, 153, 0.2)', whiteSpace: 'nowrap' } : { whiteSpace: 'nowrap' }
            },
            masterButton,
            arrowButton,
            _react2['default'].createElement(
                _materialUi.Popover,
                {
                    className: 'menuPopover',
                    open: showMenu,
                    anchorEl: anchor,
                    anchorOrigin: { horizontal: direction || 'left', vertical: 'bottom' },
                    targetOrigin: { horizontal: direction || 'left', vertical: 'top' },
                    onRequestClose: function () {
                        _this3.setState({ showMenu: false });
                    },
                    style: { marginTop: 1 },
                    useLayerForClickAway: false
                },
                _Utils2['default'].itemsToMenu(menuItems, this.menuClicked)
            )
        );
    };

    _createClass(ButtonComposed, null, [{
        key: 'propTypes',
        value: {
            buttonTitle: _react2['default'].PropTypes.oneOfType([_react2['default'].PropTypes.string, _react2['default'].PropTypes.object]).isRequired,
            masterAction: _react2['default'].PropTypes.func.isRequired,
            menuItems: _react2['default'].PropTypes.array.isRequired,
            className: _react2['default'].PropTypes.string,
            raised: _react2['default'].PropTypes.bool,
            direction: _react2['default'].PropTypes.oneOf(['left', 'right'])
        },
        enumerable: true
    }]);

    return ButtonComposed;
})(_react2['default'].Component);

exports['default'] = _MenuItemsConsumer2['default'](ButtonComposed);
module.exports = exports['default'];
