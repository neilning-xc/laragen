/**
 * Date: 2019-12-19 23:23:16
 * Author: Neil Xuchu Ning
 * Email: neil.ning@chinanetcloud.com
 * Desc: main.test.js
 */
import React from "react";
import {render, unmountComponentAtNode} from "react-dom";
import {act} from "react-dom/test-utils";

import Main from "../components/Main"

let container = null;
beforeEach(() => {
    // setup a DOM element as a render target
    container = document.createElement("div");
    document.body.appendChild(container);
});

afterEach(() => {
    // cleanup on exiting
    unmountComponentAtNode(container);
    container.remove();
    container = null;
});

it('should show hello world', () => {
    act(() => {
        render(<Main/>, container)
    });
    expect(container.textContent).toBe("Hello World")
});
