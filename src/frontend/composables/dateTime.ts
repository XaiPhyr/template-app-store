import dayjs from 'dayjs';

const formatDateTime = (
  date: string | number,
  format: string = 'DD/MM/YYYY'
) => {
  if (!date) {
    return '';
  }

  return dayjs(date).format(format);
};

const unixDateTime = (date: string | number | Date) => {
  if (!date) {
    return '';
  }

  return dayjs(date).unix();
};

export { formatDateTime, unixDateTime };
