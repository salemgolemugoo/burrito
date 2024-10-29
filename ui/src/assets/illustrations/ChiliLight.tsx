import { SVGProps } from 'react';

const ChiliLight = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={424}
    height={424}
    viewBox="0 0 424 424"
    fill="none"
    {...props}
  >
    <path
      fill="#000"
      fillRule="evenodd"
      d="M357.204 120.703c-7.074 5.154-10.593 12.475-11.653 17.541a2.5 2.5 0 0 1-4.894-1.025c1.275-6.09 5.383-14.569 13.603-20.558 8.312-6.056 20.554-9.367 37.652-5.496a2.5 2.5 0 0 1-1.104 4.877c-15.903-3.601-26.621-.426-33.604 4.661Z"
      clipRule="evenodd"
    />
    <path
      fill="#E14029"
      stroke="#000"
      strokeWidth={5}
      d="M277.392 167.459C238.626 206.53 144.707 281.21 34.109 226.461c-.684-.338-1.512.223-1.468.985 8.704 149.463 336.139 79.472 333.813-45.349-.696-37.353-23.27-47.735-43.784-45.901-18.891 1.69-31.919 17.799-45.278 31.263Z"
    />
    <mask
      id="a-chili-light"
      width={335}
      height={177}
      x={32}
      y={135}
      maskUnits="userSpaceOnUse"
      style={{
        maskType: 'alpha'
      }}
    >
      <path
        fill="#E14029"
        d="M277.392 167.459C238.626 206.53 144.707 281.21 34.109 226.461c-.684-.338-1.512.223-1.468.985 8.704 149.463 336.139 79.472 333.813-45.349-.696-37.353-23.27-47.735-43.784-45.901-18.891 1.69-31.919 17.799-45.278 31.263Z"
      />
    </mask>
    <g mask="url(#a-chili-light)">
      <path
        fill="#6BEF70"
        d="m262.162 160.256 107.409 78.61 27.241-83.28-84.059-44.365-50.591 49.035Z"
      />
      <path
        fill="#000"
        fillRule="evenodd"
        d="m312.316 108.164 87.528 46.195-28.968 88.56-112.581-82.396 54.021-52.359Zm.875 6.115-47.16 45.71 102.236 74.824 25.514-78-80.59-42.534Z"
        clipRule="evenodd"
      />
    </g>
    <path
      fill="#000"
      fillRule="evenodd"
      d="M350.377 146.358c-7.646-6.486-17.688-8.548-27.484-7.671-8.759.783-16.276 4.911-23.36 10.666-5.306 4.311-10.203 9.385-15.159 14.521-1.722 1.784-3.451 3.576-5.207 5.346-19.503 19.656-52.958 48.358-95.558 65.246-42.087 16.685-93.128 21.84-148.291-4.635 2.934 34.741 23.957 56.871 54.394 68.809 31.313 12.282 72.483 13.687 113.543 6.525 41.034-7.157 81.649-22.818 111.794-44.34 30.194-21.557 49.465-48.66 48.906-78.681-.339-18.171-5.974-29.337-13.578-35.786Zm3.234-3.813c8.986 7.621 14.986 20.323 15.343 39.505.604 32.39-20.183 60.843-51 82.845-30.867 22.037-72.219 37.936-113.84 45.196-41.597 7.255-83.772 5.934-116.228-6.796-32.6-12.787-55.496-37.173-57.74-75.704-.153-2.625 2.623-4.583 5.073-3.37 54.546 27.001 104.922 22.099 146.547 5.597 41.701-16.531 74.588-44.705 93.851-64.12 1.644-1.656 3.303-3.376 4.983-5.117 5.013-5.196 10.217-10.59 15.78-15.109 7.493-6.088 15.936-10.859 26.068-11.766 10.719-.958 22.22 1.253 31.163 8.839Z"
      clipRule="evenodd"
    />
  </svg>
);

export default ChiliLight;
